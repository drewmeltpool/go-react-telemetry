import React from  'react'
import TabletItem from './TabletItem'

function TabletList(props){

    return(
        <table className="table table-bordered table-hover">
            <tbody>
            { props.tablets.map( tablet => {
                return <TabletItem tablet = { tablet } key = {tablet.id}/>
            })}
            </tbody>
            <thead>
                <tr>
                    <th className="align-middle" scope="col">Id</th>
                    <th className="align-middle name" scope="col">Name</th>
                    <th className="align-middle" scope="col">Telemetry</th>
                </tr>
            </thead>
            
        </table> 
    )
}

export default TabletList