function App() {
  return (
    <div>
      <h1>Hola mundo</h1>
      <button onClick={async ()=>{
        const response = await fetch('http://localhost:3000/user')
        const data = await response.json()
        console.log(data)
      }}>
        Obtener datos
      </button>
    </div>
  )
}

export default App
