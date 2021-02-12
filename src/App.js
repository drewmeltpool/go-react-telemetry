import React, { useEffect, useState } from 'react'
import axios from 'axios'

import TabletList from './Tablet/TabletList'
import TabletForm from './Tablet/TabletForm'

function App() {
  const [tablets, setTablets] = React.useState([])

  useEffect(() => {
    axios.get('/tablets').then(response => {
      const tablets = response.data.sort((a,b) => a.id - b.id)
      setTablets(tablets)
    })
  }, [])

  return (
    <div className="container">
      <h1 className="h1">Tablet list</h1>
      <TabletForm/>
      <TabletList tablets = { tablets }/>

    </div>
  )
}

export default App
