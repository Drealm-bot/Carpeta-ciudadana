import axios from "axios";

type props = {
    civid: number,
    auth: string,
    filename: string,
}

export const Download = ({civid, auth, filename}: props) => {
    const url = "http://localhost:3000/repository/"+civid+"/"+filename
    const authorizationHeader = `Bearer ${auth}`;

    const handleDownload = () => {
        axios({
          method: 'GET',
          url: url,
          headers: {
            Authorization: authorizationHeader,
          },
          responseType: 'blob',
        })
          .then((response) => {
            const downloadUrl = window.URL.createObjectURL(new Blob([response.data]));
            const link = document.createElement('a');
            link.href = downloadUrl;
            link.setAttribute('download', filename);
            document.body.appendChild(link);
            link.click();
            document.body.removeChild(link);
          })
          .catch((error) => {
            console.error('Error downloading file:', error);
          });
      };

    return (
        <button onClick={handleDownload}>Descargar</button>
    );
}