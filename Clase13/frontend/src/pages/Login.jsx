import React, {useState} from 'react';
import '../styles/login.css'
import 'bootstrap/dist/css/bootstrap.min.css'


function Login() {
    const [isChecked, setIsChecked] = useState(false);
    
    const handleSubmit = (e) => {
        e.preventDefault();
    }

  return (
    <div className="form-signin">
      <div className="text-center">
        <form onSubmit={handleSubmit} className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Inicio de Sesion</h1>
          <h1 className="h3 mb-3 fw-normal">Tutorias ECYS</h1>
          <label htmlFor="inputEmail" className="visually-hidden">Usuario</label>
          <input type="text" id="userI" className="form-control" placeholder="Nombre Usuario" required
           autoFocus/>
           <br/>
           <label htmlFor="inputPassword" className="visually-hidden">Password</label>
           <input type="password" id="passI" className="form-control" placeholder="Password" aria-describedby="passwordHelpInline" //required 
            autoFocus/>
            <br />
            <div className="form-check form-switch text-left">
              <input className="form-check-input" type="checkbox" role="switch" id="flexSwitchCheckDefault"
              value={isChecked}
              onChange={e => setIsChecked(!isChecked)}
              />
              <label className="form-check-label" htmlFor="flexSwitchCheckDefault">Tutor</label>
            </div>
            <br />
            <button className="w-100 btn btn-lg btn-primary" type="submit">Iniciar Sesion</button>
            <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
            <br/>
        </form>
      </div>
    </div>
  )
}

export default Login