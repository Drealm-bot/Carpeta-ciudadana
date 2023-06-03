import Axios from "axios";
import { useState } from "react";

const signupForm = {
    "civid": "",
    "name": "",
    "address": "",
    "email": ""
}

export const Signup = () => {
    const url = "http://localhost:3000/signup"
    const [formData, setFormData] = useState(signupForm);
    const { civid, name, address, email } = formData;

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
                civid: parseInt(formData.civid),
                name: formData.name,
                address: formData.address,
                email: formData.email
            }, {headers: {
                'Content-Type': 'application/json; charset=UTF-8'
              }
            }).then(resp => {
                console.log(resp.data);
            })
            setFormData(signupForm);
        };
        return (
            <div>
                <h1>Registro</h1>
                <p>¡Regístrate si todavía no tienes una carpeta ciudadana!</p>

                <form method="dialog" onSubmit={onSubmit}>
                    <table>
                        <tbody>
                        <tr>
                            <th><label htmlFor="name">Nombre</label></th>
                                <th><input
                                    type="text"
                                    name="name"
                                    value={name}
                                    onChange={onChange}
                                /></th>
                                <th>                       </th>
                        </tr>
                    

                        <tr><th><label htmlFor="civid">Identificación</label></th>
                        <th><input
                            type="text"
                            name="civid"
                            value={civid}
                            onChange={onChange}
                        /></th></tr>

                        <tr><th><label htmlFor="address">Dirección</label></th>
                        <th><input
                            type="text"
                            name="address"
                            value={address}
                            onChange={onChange}
                        /></th></tr>

                        <tr><th><label htmlFor="email">Email</label></th>
                        <th><input
                            type="email"
                            name="email"
                            value={email}
                            onChange={onChange}
                        /></th></tr></tbody>
                    </table>
                    <button type="submit">Registrarte</button>
                </form>
            </div>
            
        )
        
}