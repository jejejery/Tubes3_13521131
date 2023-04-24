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
    handleAlgoChange = (event) =>{
      this.setState({algo: event.target.value})
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
          <SideBar style={{ margin: '0 10px' }} handleAlgoChange= {this.handleAlgoChange} algo = {this.state.algo}/>
          <div className={styles.chatbox}>
            <h1>to be continued...</h1>
          </div>
          <div style={{ margin: '0 10px', flex: 1, position: 'relative' }}>
            <Form onSubmit = {this.handleSubmit}>
            {/* <Form.Group controlId="formMessage"> */}
            <div style = {{display: 'flex', alignItems: 'flex-end', position: 'absolute', bottom: '2%', width: '100%'}}>
            <Form.Control
            type="text"
            placeholder="Send a message..."
            onChange={this.handleReadText}
            style={{ fontSize: '24px', marginLeft: '5.5%', marginRight: '0px', width: '80%'}}
            />
            <Button type = "submit" style = {{marginLeft: '5px', height: '50px'}}> send </Button>
            </div>
            {/* </Form.Group> */}
            </Form>
         </div>
       </div>
      </div>
            );
     }
}
     
     
  export default MainGPT;