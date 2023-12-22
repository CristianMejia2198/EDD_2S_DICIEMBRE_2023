import React, { useState } from 'react'
import 'bootstrap/dist/css/bootstrap.min.css'

function Home() {
  const [numero, setNumero] = useState(0)

  const accion = async() =>{
    console.log(numero)
    const response = await fetch("http://localhost:4000/ingresar-arbol-b",{
      method: 'POST',
      headers: {
        "Content-Type":"application/json"
      },
      body: JSON.stringify({
        valor: numero
      })
    })
    const result = await response.json()
    console.log(result)
  }

  const report = async() => {
    const response = await fetch('http://localhost:4000/reporte', {
      method: 'POST',
      headers: {
        "Content-Type":"application/json"
      },
      body: JSON.stringify({
        nombre: "arbolitoF",
        estructura: "Arbol B"
      })
    })
    console.log(response)
    const result = await response.json()
    console.log(result)
  }

  return (
    <div>
      <h1>HOME</h1>
      <div className="container text-center">
        <div className="row">
          <div className="col align-self-center">
            <input type="number" onChange={ e => setNumero(parseInt(e.target.value))} placeholder='15'/>
          </div>
        </div>
        <div className="row">
          <div className="col align-self-center">
            <button type="button" className="btn btn-outline-success" onClick={accion}>Enviar</button>
            <button type="button" className="btn btn-outline-success" onClick={report}>Reporte</button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home
