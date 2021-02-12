import React from 'react'
import TelemetryItem from "./TelemetryItem";


function TabletTelemetry(props) {
    return(
        <div className="d-flex flex-wrap">
            { props.telemetries.map((telemetry, index) => {
                return <TelemetryItem telemetry = {telemetry} key = {index}/>  }) }
        </div>
    )
}

export default TabletTelemetry