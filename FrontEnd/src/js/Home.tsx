import {SyntheticEvent, useState} from 'react';
import '../css/Header.css'
import {useNavigate} from "react-router-dom";
import { MDBInput } from 'mdb-react-ui-kit';
import {

    MDBBtn
  } from 'mdb-react-ui-kit';

const Home = (props: { name: string }) => {

    const [loc, setLoc] = useState('');
    const [name, setName] = useState('');
    const [temp, setTemp] = useState('');
    const [cond, setCond] = useState('');
    const [img, setImg] = useState('');


    const navigate = useNavigate();

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        const response = await fetch('http://localhost:8000/api/weather', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
            body: JSON.stringify({
                loc
            })
        });

        var content = await response.json();
        setName(content.location.name)
        setTemp(content.current.temp_c)
        setCond(content.current.condition.text)
        setImg(content.current.condition.icon)

    }

        
    
    
    

    let result;

    let menu;

    if (props.name === '') {
        menu = (
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
    )
        }else{
            menu = (
                <div className='searchbar'>
                <h2> Please, write in your city </h2>
                <form onSubmit={submit}>
                <div className='search'>
                <MDBInput onChange={e => setLoc(e.target.value)} label='Form control lg' id='formControlLg' type='text' size='lg' />
                <MDBBtn type='submit'>Primary</MDBBtn>
                </div>
                </form>
                </div>
            )
        }

        if(name !== ''){
            result=(
                <div className='result'>
            <img src={img} />
            <p>Condition: {cond} </p>
            <p>Temperature: {temp}</p>
            <p>Name: {name}</p>

        </div>
            )
        }

    return (
        <div>
        {menu}
        {result}
        </div>
    );
};


export default Home;