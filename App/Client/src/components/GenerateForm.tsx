import Axios from "axios";
import { useState } from "react";

const generateForm = {
    "civid": "",
    "email": ""
}

export const Generate = () => {
    const url = "http://localhost:3000/generate"
    const [formData, setFormData] = useState(generateForm);
    const { civid, email } = formData;

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
                email: formData.email
            }, {headers: {
                'Content-Type': 'application/json; charset=UTF-8'
              }
            }).then(resp => {
                console.log(resp.data);
            })
            setFormData(generateForm);
        };
        return (
            <div>
                <h1>Generar Contraseña</h1>
                <p>Ingrese su información de usuario.</p> 

                <form onSubmit={onSubmit}>
                <table>
                       <tbody>
                    <tr><th><label htmlFor="civid">Identificación</label></th>
                    <th><input
                        type="text"
                        name="civid"
                        value={civid}
                        onChange={onChange}
                    /></th></tr>
                    
                    <tr><th><label htmlFor="email">Email</label></th>
                    <th><input
                        type="email"
                        name="email"
                        value={email}
                        onChange={onChange}
                    /></th></tr>
                    </tbody> 
                </table>

                    <button type="submit">Generar</button>
                    
                </form>
                
            </div>
        )
}