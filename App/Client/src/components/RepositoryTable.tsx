import Axios from "axios"
import { Download } from "./DownloadFile"
import { useEffect, useState } from "react"
import { Authenticate } from "./AuthenticateFile";

interface Archive {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    Owner: number;
    FullName: string;
    Name: string;
    Type: string;
    Path: string;
    IsAuthenticated: boolean;
  }

type props = {
    civid: number,
    auth: string,
    uploadedFile: any
}

export const Repository = ({civid, auth, uploadedFile}: props) => {
    const url = "http://localhost:3000/repository/"+ civid
    console.log(url)

    const [archives, setArchives] = useState<Archive[]>([])
    const [refresh, setRefresh] = useState(false);

    const config = {
        headers: { Authorization: `Bearer ${auth}` }
    };

    useEffect(() => {
        Axios.get(url, config).then(res => {
            console.log(res.data)
            setArchives(res.data)
        })
    },[uploadedFile, auth, refresh])

    const arr = archives.map((data, index) => {
        return (
          <tr key={index}>
            <td>{data.Name}</td>
            <td>{data.Type}</td>
            <td><Download civid={civid} auth={auth} filename={data.FullName} /></td>
            <td><Authenticate civid={civid} auth={auth} filename={data.FullName} onAuthenticate={() => setRefresh(!refresh)}/></td>
            <td>{data.IsAuthenticated ? "si":"no"}</td>
          </tr>
        );
      });
    if (archives.length === 0) {
        return(
            <h2>Sube un documento.</h2>
        )
    }else{
        return(
            <table>
            <thead>
              <tr>
                <th>Nombre</th>
                <th>Tipo</th>
                <th>Descargar</th>
                <th>Autenticar</th>
                <th>Autenticado</th>
              </tr>
            </thead>
            <tbody>
                {arr}
            </tbody>
            </table>
        )
    }
    
}