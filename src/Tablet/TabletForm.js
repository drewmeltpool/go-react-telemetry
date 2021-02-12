import React from 'react';
import axios from 'axios'

const FormData = () => {
    const data = {}
    const temetryId = document.querySelector('#telemetry_id')
    data.id = temetryId.value
    const telemetryBattery = document.querySelector('#telemetry_battery')
    data.battery = telemetryBattery.value
    const telemetryCurrentVideo = document.querySelector('#telemetry_currentVideo')
    data.currentVideo = telemetryCurrentVideo.value
    
    const date = new Date();
    const formatted = date.toISOString();

    axios.post('/tablets', {
        id: Number(data.id),
        battery: Number(data.battery),
        currentVideo: String(data.currentVideo),
        deviceTime: formatted
      })
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
}



function TabletForm() {
    return(  
        <div className="position-fixed">
            <p className="bg-primary w-100 text-uppercase fs-6 text-white text-center rounded-top" onClick={FormData}>Form</p>
            <form className="bg-dark p-2 pr-4 d-flex flex-column">
                <div className="input-group d-flex align-items-center m-2">
                    <span className="input-group-text p-1 rounded-0"><ion-icon name="key-outline"></ion-icon></span>
                    <input type="number" min="0" className="form-control h-100 d-flex" id="telemetry_id" placeholder="TelemetryId"/>
                </div>
                <div className="input-group d-flex align-items-center m-2">
                    <span className="input-group-text p-1 rounded-0" id="addon-wrapping"><ion-icon name="battery-full-outline"></ion-icon></span>
                    <input type="number" min="0" max="100" className="form-control h-100 d-flex" id="telemetry_battery" placeholder="Battery"/>
                </div>
                <div className="input-group d-flex align-items-center m-2">
                    <span className="input-group-text p-1 rounded-0" id="addon-wrapping"><ion-icon name="videocam-outline"></ion-icon></span>
                    <input type="text" className="form-control h-100 d-flex" id="telemetry_currentVideo" placeholder="Video"/>
                </div>
            </form>
            <a href="http://localhost:3000" className="btn btn-primary w-100 btn__form" onClick={FormData}>Update Tablet</a>
        </div>
    )
}

export default TabletForm