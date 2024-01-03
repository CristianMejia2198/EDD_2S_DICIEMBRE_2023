import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import Login from "./pages/Login";
import "./App.css";
import Principal_Admin from "./pages/Administrador/Principal_Admin";
import Carga_Archivos from "./pages/Administrador/Carga_Archivos";
import Libros_Control from "./pages/Administrador/Libros_Control";
import Tabla_Alumnos from "./pages/Administrador/Tabla_Alumnos";
import Cargar_Libro from "./pages/Tutor/Cargar_Libro";
import Carga_Publicacion from "./pages/Tutor/Carga_Publicacion";
import Principal_Estudiante from "./pages/Estudiante/Principal_Estudiante";
import Reporte from "./pages/Administrador/Reporte";

function App() {
  const [count, setCount] = useState(0);

  return (
    /*<>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/principal/estudiante" element={<Estudiante />} />
        <Route path="/principal/tutor" element={<Tutor />} />
        <Route path="/principal/admin" element={<PrinicipalAdmin />} />
        <Route path="/principal/admin/alumnos" element={<TablaAlumnos />} />
        <Route path="/principal/tutor/libro" element={<LibrosAdmin />} />

        <Route path="/principal/admin/cargar-libro" element={<LibrosTutor />} />
        <Route
          path="/principal/admin/carga-archivos"
          element={<Administrador />}
        />
      </Routes>
    </>*/
    <>
      <Routes>
        <Route path="/" element={<Login />} />

        <Route path="/principal/admin" element={<Principal_Admin />} />
        <Route path="/principal/admin/archivo" element={<Carga_Archivos />} />
        <Route path="/principal/admin/libros" element={<Libros_Control />} />
        <Route path="/principal/admin/alumnos" element={<Tabla_Alumnos />} />
        <Route path="/principal/admin/reporte" element={<Reporte />} />

        <Route path="/principal/tutor" element={<Cargar_Libro />} />
        <Route
          path="/principal/tutor/publicacion"
          element={<Carga_Publicacion />}
        />

        <Route
          path="/principal/estudiante"
          element={<Principal_Estudiante />}
        />
      </Routes>
    </>
  );
}

export default App;
