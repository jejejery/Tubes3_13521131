import React, {useEffect} from 'react';
import styles from './MainGPT.module.css';
import SideBar from '../../components/SideBar/SideBar'; 
import { Form, Button } from "react-bootstrap";
import axios from 'axios';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.css';
// import { Route } from 'react-router';



class MainGPT extends React.Component {  
    constructor(props){
      super(props);
      this.state = {
        value : ""
      };
      
    }
    handleReadText = (event) =>{
      this.setState({value: event.target.value})
      console.log(this.state.value)
    }
    handleSubmit = (event) => {
      event.preventDefault()
      
    }

    render() { 
      return (
      <div className={styles.main}>
        <div style={{ display: 'flex', flex: '1 auto', justifyContent: 'center', position: 'relative' }}>
          <SideBar style={{ margin: '0 10px' }} />
          <div className={styles.chatbox}>
            <h1>to be continued...</h1>
          </div>
          <div style={{ margin: '0 10px', flex: 1, position: 'relative' }}>
            <Form.Control
            type="text"
            placeholder="Send a message..."
            onChange={this.handleReadText}
            style={{ fontSize: '24px', position: 'absolute', bottom: '2%', margin: '0 10px', marginLeft: 'auto' }}
            />
         </div>
       </div>
      </div>
            );
     }
}
     
     
  export default MainGPT;