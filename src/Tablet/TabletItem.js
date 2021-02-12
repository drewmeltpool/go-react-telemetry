import React from 'react'
import TabletTelemetry from './Telemetry'

function TabletItem({tablet}) {
    return (
        <tr className="">
            <th className="align-middle" scope="row">
                {tablet.id}
            </th>
            <td className="align-middle">
                {tablet.name}
            </td>
            <td className="align-middle">
                <TabletTelemetry telemetries = {tablet.telemetry}/>
            </td>
        </tr>

    )
}

export default TabletItem