import React from 'react'

const telemetryVideo = value => String(value) 

function Telemetryitem({telemetry}) {
    return(

        <div className="dropdown m-1">
            <button className="btn btn-dark dropdown-toggle d-flex align-items-center" type="button" id="dropdownMenuButton" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                <span className="tablet-id font-weight-bold">{telemetry.id}</span>
                <ion-icon name="tablet-portrait"></ion-icon>
            </button>
            <ul className="dropdown-menu mt-2 p-0" aria-labelledby="dropdownMenuButton">
                <li className="list-group-item d-flex flex-column align-items-start m-0 p-0 bg-blue-100">
                    <p className="d-flex align-items-center w-100 bg-primary text-white">
                        <ion-icon name="battery-full-outline"></ion-icon>
                        <span className="font-weight-bold">{telemetry.battery}</span>
                    </p>
                    <p className="d-flex align-items-center w-100">
                        <ion-icon name="time-outline"></ion-icon>
                        <span className="font-weight-bold">{telemetry.deviceTime}</span>
                    </p>
                    <p className="d-flex align-items-center w-100">
                        <ion-icon name="server-outline"></ion-icon>
                        <span className="font-weight-bold"> {telemetry.timeStamp}</span>
                    </p>
                    <p className="d-flex align-items-center w-100">
                        <ion-icon name="videocam-outline"></ion-icon>
                        <a href={telemetryVideo(telemetry.currentVideo)} target="__blank" className="font-weight-bold"> {telemetryVideo(telemetry.currentVideo)}</a>
                    </p>
                </li>  
            </ul>
        </div> 
          
    )
}
export default Telemetryitem