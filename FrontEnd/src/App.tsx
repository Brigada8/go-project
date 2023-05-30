import React, {useEffect, useState} from 'react';
import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Header from "./js/Header"
import Home from "./js/Home"
import Login from "./js/Login"
import Register from "./js/Register"


function App() {

  const [name, setName] = useState('');

  useEffect(() => {
      (
          async () => {
              const response = await fetch('http://localhost:8000/api/user', {
                  headers: {'Content-Type': 'application/json'},
                  credentials: 'include',
              });

              const content = await response.json();

              setName(content.Name);
              console.log(content)
          }
      )();
  });

  return (
    <>
                  <BrowserRouter>
                  <Header  name={name} setName={setName}/>

                <main className="form-signin">
                    <Routes>
                    <Route path="/" element={<Home name={name}/>}/>
                    <Route path="/login" element={<Login setName={setName}/>}/>
                    <Route path="/register" element={<Register/>}/>
                    </Routes>
                </main>
            </BrowserRouter>
    </>
  )
}

export default App
