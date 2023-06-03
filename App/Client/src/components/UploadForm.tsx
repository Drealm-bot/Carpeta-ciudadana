import Axios from "axios"
import { useState } from "react"


type props = {
    civid: number,
    auth: string,
    onFileUpload: any
}

export const Upload = ({civid, auth, onFileUpload}: props) => {
    const [file, setFile] = useState<File | undefined>(undefined)

    const url = "http://localhost:3000/repository/" + civid + "/upload"
    const config = {
        headers: { 
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${auth}`,
     }
    };

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log(file);
        Axios.post(url, {
            file: file
        }, config
        ).then(resp => {
            console.log(resp.data);
            onFileUpload(file);
        })
    };

    return(
        <div>
            <h1>Upload single file with fields</h1>

            <form onSubmit={onSubmit}>
            <p><label htmlFor="file">File</label>
                    <input
                        type="file"
                        name="file"
                        onChange={(e) => {
                            const files = e.target.files;
                            if (files && files.length > 0) {
                              setFile(files[0]);
                            }
                          }}
                    /></p>
                    <input type="submit" value="Submit"></input>
            </form>
        </div>
    )

}