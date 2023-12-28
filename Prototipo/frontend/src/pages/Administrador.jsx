import React from "react";

function Administrador() {
  const reportes = (e) => {
    e.preventDefault();
  };

  const uploadFile = (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = (event) => {
      const content = event.target.result;
      console.log("Contenido del archivo CSV:", content);
      const parsedData = parseCSV(content);
      console.log("Datos analizados:", parsedData);
    };

    reader.onerror = (error) => {
      console.error("Error al leer el archivo:", error);
    };

    reader.readAsText(file);
  };

  const parseCSV = (csvContent) => {
    const rows = csvContent.split("\n");
    const encabezado = rows.slice(1);
    const sinRetorno = encabezado.map((row) => row.trim());
    const data = sinRetorno.map((row) => row.split(","));
    return data;
  };

  const onChange1 = (e) => {};

  const salir = (e) => {
    e.preventDefault();
    console.log("Listo");
    window.open("/", "_self");
  };

  return (
    <div className="form-signin1">
      <div className="text-center">
        <form className="card card-body">
          <h1 className="h3 mb-3 fw-normal">Administrador</h1>
          <br />
          <h4 className="h3 mb-3 fw-normal">Cargar Archivos</h4>
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar csv</label>
            <input
              className="form-control"
              id="inputGroupFile01"
              type="file"
              accept=".csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet, application/vnd.ms-excel"
              onChange={uploadFile}
            />
          </div>
          <br />
          <div className="input-group mb-3">
            <label className="input-group-text">Cargar json</label>
            <input
              className="form-control"
              id="inputGroupFile02"
              type="file"
              accept="application/json"
              onChange={onChange1}
            />
          </div>
          <br />
          <center>
            <button className="w-50 btn btn-outline-primary" onClick={reportes}>
              Reportes
            </button>
          </center>
          <br />
          <center>
            <button className="w-50 btn btn-outline-success" onClick={salir}>
              Salir
            </button>
          </center>
          <p className="mt-5 mb-3 text-muted">EDD 201700918</p>
          <br />
        </form>
      </div>
    </div>
  );
}

export default Administrador;
