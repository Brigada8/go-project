import {useEffect, useState} from 'react';
import './App.css'
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Header from "./js/Header"
import Home from "./js/Home"
import Login from "./js/Login"
import Register from "./js/Register"
import Footer from "./js/Footer"


function App() {

  const [name, setName] = useState('');

  useEffect(() => {
      (
          async () => {
              const response = await fetch('https://goback-1ssr.onrender.com/api/user', {
                mode: 'cors',
                headers: {
                    'Access-Control-Allow-Origin':'*',
                    'Content-Type': 'application/json'
                  },
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

                <main className="">
                  
                    <Routes>
                    <Route path="/" element={<Home name={name}/>}/>
                    <Route path="/login" element={<Login setName={setName}/>}/>
                    <Route path="/register" element={<Register/>}/>
                    </Routes>
                </main>
            </BrowserRouter>
            <Footer/>

    </>
  )
}

export default App
