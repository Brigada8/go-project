import React from 'react';

const Home = (props: { name: string }) => {
    return (
        <div>
            {props.name ? 'Hi ' + props.name : 'Log in to use our service'}
        </div>
    );
};

export default Home;