import Axios from "axios";
import { useEffect, useState } from "react";

const loginForm = {
    "civid": "",
    "password": ""
}

const authDefaults = {
    "id": 0,
    "message": "",
    "token": "",
}

type props = {
    extractAuth: (data: [number, string]) => void
}

export const Login = ({ extractAuth }: props) => {
    const url = "http://localhost:3000/login"
    const [formData, setFormData] = useState(loginForm);
    const [auth, setAuth] = useState(authDefaults);
    const { civid, password } = formData;

        const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
            setFormData((prevState) => ({
              ...prevState,
                [e.target.name]: e.target.value
            }));
        }
    
        const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
            e.preventDefault();
            console.log(formData);
            Axios.post(url, {
                id: parseInt(formData.civid),
                password: formData.password
            }, {headers: {
                'Content-Type': 'application/json; charset=UTF-8'
              }
            }).then(resp => {
                console.log(resp.data);
                setAuth(resp.data);
            })
            setFormData(loginForm);
        };

        useEffect(() => {
            console.log(auth)
            extractAuth([auth.id, auth.token])
        }, [auth]);



        return (
            <div>
                <h1>Login</h1>
                <p>Ingrese la contraseña que le llegó a su correo.</p> 

                <form onSubmit={onSubmit}>
                    <table>
                        <tbody>
                            <tr>
                    <th><label htmlFor="civid">Identificación</label></th>
                    <th><input
                        type="text"
                        name="civid"
                        value={civid}
                        onChange={onChange}
                    /></th></tr>
                    
                    <tr><th><label htmlFor="password">Contraseña</label></th>
                    <th><input
                        type="password"
                        name="password"
                        value={password}
                        onChange={onChange}
                    /></th></tr>
                    </tbody>
                    </table>

                    <button type="submit">Ingresar</button>
                </form>
            </div>
        )
}