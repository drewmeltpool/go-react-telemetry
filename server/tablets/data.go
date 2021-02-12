package tablets

import (
	"database/sql"
	"time"
	"encoding/json"
)

//NullString ...
type NullString struct {
	sql.NullString
}

//MarshalJSON ...
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

//UnmarshalJSON ...
func (s *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		s.String, s.Valid = "", false
		return nil
	}
	s.String, s.Valid = string(data), true
	return nil
}

// Tablet ...
type Tablet struct {
	ID			int64		`json:"id"`
	Name        string		`json:"name"`
	Telemetry	[]*Device	`json:"telemetry"`
}

//Device ...
type Device struct {
	ID				int64		`json:"id"`
	Battery			int64		`json:"battery"`
	DeviceTime		string		`json:"deviceTime"`
	TimeStamp		string		`json:"timeStamp"`
	CurremtVideo	NullString	`json:"currentVideo"`
}

//UpdateDev ...
type UpdateDev struct {
	ID				int64	`json:"id"`
	Battery			int64	`json:"battery"`
	CurrentVideo	string	`json:"currentVideo"`
	DeviceTime		string	`json:"deviceTime"`
}

//Tablets ...
type Tablets struct {
	TabletrsArr []*Tablet `json:"Tablets"`
}

//Devices ...
type Devices struct {
	Devs []*Device `json:"Devices"`
}

//Store ...
type Store struct {
	Db *sql.DB
}

//NewStore ...
func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

//ListOfTablets returns a list of all tabletss
func (s *Store) ListOfTablets() ([]*Tablet, error) {
	rows, err := s.Db.Query("SELECT id,name FROM tablet ORDER BY id DESC LIMIT 50")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Tablet
	for rows.Next() {
		var b Tablet
		if err := rows.Scan(&b.ID,&b.Name); err != nil {
			return nil, err
		}
		res = append(res, &b)
	}

	var fullTablets []*Tablet
	if res == nil {
		fullTablets = make([]*Tablet, 0)
	} else {
		for i := 0; i < len(res); i++ {
			machines, err := s.GetTelemetryByID(res[i].ID)
			if err != nil {
				return nil, err
			}
			fullBalancer := Tablet{
				ID:			res[i].ID,
				Name:		res[i].Name,
				Telemetry:	machines,
			}
			fullTablets = append(fullTablets, &fullBalancer)
		}

	}
	result := &Tablets{fullTablets}

	return result.TabletrsArr, err
}

//GetTelemetryByID ...
func (s *Store) GetTelemetryByID(id int64) ([]*Device, error) {
	rows, err := s.Db.Query(`select tele.id, tele.battery, tele.deviceTime, tele.timeStamp, tele.currentVideo from tablet tab
	join telemetry tele on tabletId = tab.id
	where tab.id = ?`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Device
	for rows.Next() {
		var b Device
		if err := rows.Scan(&b.ID,&b.Battery,&b.DeviceTime,&b.TimeStamp,&b.CurremtVideo); err != nil {
			return nil, err
		}
		res = append(res, &b)
	}

	var fullDevices []*Device
	if res == nil {
		fullDevices = make([]*Device, 0)
	} else {
		for i := 0; i < len(res); i++ {
			fullDevice := Device{
				ID:        		res[i].ID,
				Battery:        res[i].Battery,
				DeviceTime:     res[i].DeviceTime,
				TimeStamp:		res[i].TimeStamp,
				CurremtVideo:	res[i].CurremtVideo,
			}
			fullDevices = append(fullDevices , &fullDevice)
		}

	}

	result := &Devices{fullDevices}
	return result.Devs, err
}

//UpdateDevice updates a device in DB
func (s *Store) UpdateDevice(id int64, Battery int64, CurrentVideo string, DeviceTime string) error {
	t := time.Now()
	TimeStamp := t.Format("2006-01-02T15:04:05.000Z")
	_, err := s.Db.Exec("update telemetry set battery=?, currentVideo=?, timeStamp=?, deviceTime=? where id=?", Battery, CurrentVideo, TimeStamp, DeviceTime, id)
	return err
}
