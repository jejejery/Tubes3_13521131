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
        text : "",
        algo : true
      };
      
    }
    handleReadText = (event) =>{
      this.setState({text: event.target.value})
      // console.log(this.state.value)
    }
    handleSubmit = async (event) => {
      event.preventDefault()
      console.log(this.state.text)
      const data = {
        InputText : this.state.text,
        Algorithm: this.state.algo,
      };
      try{
        console.log(data)
        const res = await axios.post("http://localhost:8000/api/input", data)
        console.log(res.data)
      }catch(error){
        console.log(error)
      }
      


      
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
            <Form onSubmit = {this.handleSubmit}>
            <Form.Group controlId="formMessage">
            <Form.Control
            type="text"
            placeholder="Send a message..."
            onChange={this.handleReadText}
            style={{ fontSize: '24px', position: 'absolute', bottom: '2%', margin: '0 10px', marginLeft: 'auto' }}
            />
            <Button type = "submit" style = {{position: 'absolute', right: '0', bottom: '0'}}> send </Button>
            </Form.Group>
            </Form>
         </div>
       </div>
      </div>
            );
     }
}
     
     
  export default MainGPT;