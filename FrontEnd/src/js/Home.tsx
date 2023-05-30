import React from 'react';
import { MDBBtn } from 'mdb-react-ui-kit';
import '../css/Header.css'
import {useNavigate} from "react-router-dom";

const Home = (props: { name: string }) => {
    const navigate = useNavigate();
    return (
        <div className='back'>
            <div className='p-5 text-center bg-image' style={{ backgroundImage: "url('https://wallpaperaccess.com/full/1540016.jpg')", height: '500px', width: "100%" }}>
        <div className='mask' style={{ backgroundColor: 'rgba(0, 0, 0, 0.6)' }}>
          <div className='d-flex justify-content-center align-items-center h-100'>
            <div className='text-white'>
              <h1 className='mb-3'>You need to Log In</h1>
              <h4 className='mb-3'>To use our service</h4>
              <MDBBtn onClick={()=>navigate("/login")} tag="a" outline size="lg">
                Log In
              </MDBBtn>
            </div>
          </div>
        </div>
      </div>
        </div>
    );
};

export default Home;