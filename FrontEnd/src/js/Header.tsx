import '../css/Header.css'

import { MDBBtn } from 'mdb-react-ui-kit';
import {useNavigate} from "react-router-dom";


const Header = (props: { name: string, setName: (name: string) => void }) => {

      const logout = async () => {
        await fetch('https://goback-1ssr.onrender.com/api/logout', {
            headers: {'Content-Type': 'application/json'},
            method: 'GET',
            credentials: 'include',
        });

        props.setName('');
        window.location.reload();
    }

    const navigate = useNavigate();


    let menu;

    if (props.name === undefined) {
        menu = (
          <div className='define'>
          <MDBBtn onClick={()=>navigate("/login")}>LogIn</MDBBtn>
          <MDBBtn onClick={()=>navigate("/register")} >Register</MDBBtn>
          </div>
        )
    }else{
      menu = (
        <div className='define'>
          <MDBBtn onClick={logout} >LogOut</MDBBtn>
          </div>
      )
    }

    return (
        <>
        <div className="header">
            <div className='define'>
            <a onClick={()=>navigate("/")}> <img className='logo' src='https://www.transparentpng.com/thumb/temperature/climate-control-home-temperature-png-17.png'/> </a>
            <h1> Get your current weather now! </h1>
            </div>
            {menu}
        </div>
      </>
    );
  }
  export default Header