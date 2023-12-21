import { useState } from 'react'
import './App.css'
import {Route, Routes} from 'react-router-dom'
import Home from './pages/Home'
import Inicio from './pages/Inicio'
import Login from './pages/Login'

function App() {
  const [count, setCount] = useState(0)

  return (
      <>
        <Routes>
          <Route path = '/' element={<Home/>} />
          <Route path = '/inicio' element={<Inicio/>} />
          <Route path = '/login' element={<Login/>} />
        </Routes>
      </>
  )
}

export default App
