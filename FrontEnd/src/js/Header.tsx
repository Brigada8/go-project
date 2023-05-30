import '../css/Header.css'
import { MDBInput } from 'mdb-react-ui-kit';
import { MDBBtn } from 'mdb-react-ui-kit';


const Header = (props: { name: string, setName: (name: string) => void }) => {

      const logout = async () => {
        await fetch('http://localhost:8000/api/logout', {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        });

        props.setName('');
    }

    return (
        <>
        <div className="header">
            <div className='define'>
            <img className='logo' src='https://www.transparentpng.com/thumb/temperature/climate-control-home-temperature-png-17.png'/>
            <h1> Get your current weather now! </h1>
            </div>
            <div className='define'>
            <input className='text' placeholder='text' />
            <MDBBtn onClick={logout} >Button</MDBBtn>
            </div>
        </div>
      </>
    );
  }
  export default Header
  