import { useEffect, useState } from "react"
import { Generate } from "./components/GenerateForm"
import { Login } from "./components/LoginForm"
import { NavBar } from "./components/NavBar"
import { Signup } from "./components/SignupForm"
import 'bootstrap/dist/css/bootstrap.min.css'
import { Repository } from "./components/RepositoryTable"
import { Upload } from "./components/UploadForm"

function App() {
  const [auth, setAuth] = useState<[number, string]>([0, ""])
  const getAuth = (data: [number, string]) => {
    setAuth(data)
    console.log(data)
  }
    const [uploadedFile, setUploadedFile] = useState<File | undefined>(undefined);
  
    const handleFileUpload = (file: File) => {
      setUploadedFile(file);
    };

  return (
    <div className="App">
      <NavBar />
      <Signup />
      <Generate />
      <Login extractAuth={getAuth}/>
      <Repository civid={auth[0]} auth={auth[1]} uploadedFile={uploadedFile}/>
      <Upload civid={auth[0]} auth={auth[1]} onFileUpload={handleFileUpload}/>
    </div>
  )
  
  //
  //
  //
  //
}

export default App
