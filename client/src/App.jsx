import './App.css'

function App() {

  return (
    <div className="App">
        <h1>HOLA MUNDO</h1>
        <button onClick={async () => {
          const response = await fetch("http://localhost:3000/users")
          const data = await response.json()
          console.log(data)
        }}>Obtener datos del back</button>
    </div>
  )
}

export default App
