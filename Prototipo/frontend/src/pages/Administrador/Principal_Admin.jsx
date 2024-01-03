import React from "react";

function Principal_Admin() {
  const archivos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/archivo", "_self");
  };

  const alumnos = (e) => {
    e.preventDefault();
    window.open("/principal/admin/alumnos", "_self");
  };

  const libros = (e) => {
    e.preventDefault();
    window.open("/principal/admin/libros", "_self");
  };

  const reporte = (e) => {
    e.preventDefault();
    window.open("/principal/admin/reporte", "_self");
  };

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    localStorage.clear();
    window.open("/", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Administrador - Principal</h1>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={archivos}>
              Carga Archivos
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={alumnos}>
              Ver Alumnos
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={libros}>
              Ver Libros
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={reporte}>
              Reportes
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Salir
            </button>
          </center>
        </form>
      </div>
    </div>
  );
}

export default Principal_Admin;
