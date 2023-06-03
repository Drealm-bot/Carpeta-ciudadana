import axios from "axios";

type props = {
    civid: number,
    auth: string,
    filename: string,
    onAuthenticate: () => void;
}

export const Authenticate = ({civid, auth, filename, onAuthenticate}: props) => {
    const url = "http://localhost:3000/repository/"+civid+"/authenticate/"+filename
    const authorizationHeader = `Bearer ${auth}`;

    const handleAuthenticate = () => {
        axios({
          method: 'GET',
          url: url,
          headers: {
            Authorization: authorizationHeader,
          }}).then(() => {
            onAuthenticate();
          })
    }

    return (
        <button onClick={handleAuthenticate}>Autenticar</button>
    );
}