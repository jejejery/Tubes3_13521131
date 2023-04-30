import React, {useEffect} from 'react';
import styles from './MainGPT.module.css';
import SideBar from '../../components/SideBar/SideBar'; 
import { Form, Button } from "react-bootstrap";
import axios from 'axios';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.css';
import { SpeechBubble} from 'react-kawaii';
// import { Route } from 'react-router';



class MainGPT extends React.Component {  
    constructor(props){
      super(props);
      this.myQuestion = React.createRef(); // membuat ref
      this.state = {        
        algo : 0,//algo code: 0 is null, 1 is KMP, and 2 is BM
        qaBlocks : []
      };
      
    }
    handleAlgoChange = (event) =>{
      this.setState({algo: event.target.value})
    }
    handleReadText = (event) =>{
      this.myQuestion.current.value = event.target.value;
      
    }
    handleSubmit = async (event) => {
      event.preventDefault()
      console.log(this.myQuestion.current.value)
      console.log(this.state.algo)


      const data = {
        InputText :this.myQuestion.current.value,
        Algorithm: this.state.algo,
      };
      

      //insert into array
      let temp = this.state.qaBlocks
      temp.push(this.render_q_block(this.myQuestion.current.value,"belum adaaaa hehe :v"))
      this.setState({qaBlocks: temp})
      
      //reset the question
      this.myQuestion.current.value = ""
      //try-catch
      try{
        console.log(data)
        const res = await axios.post("http://localhost:8000/api/input", data)
        console.log(res.data)

      // mengubah nilai input menjadi kosong (empty string)
       
      }catch(error){
        console.log(error)
      }
    }
    
    render_q_block(question,answer){
      return (
        <div>
          <div class = {styles.qblock}>
            <div >
              <SpeechBubble size={80} mood="excited" color="rgb(90, 88, 177)" /> 
              <span>{question}</span>
            </div>
          </div>

          <div class = {styles.ablock}>
            <div>
              <SpeechBubble size={80} mood="happy" color="#7e4474" /> 
              <span>{answer}</span>
            </div>          
          </div>
        </div>
        
        
      )
    }

    render() { 
      //To Do in here: read recent qna from database to show em up in window


      return (
      
      <div className={styles.main}>
        <div style={styles.first}>
          <SideBar style={{ margin: '0 10px' }} handleAlgoChange= {this.handleAlgoChange} />
          <div className={styles.chatbox}>
            {this.state.qaBlocks}
          </div>
          <div style={{ left: '30%', bottom: '3%', flex: 5, position: 'absolute', borderColor:'black'}}>
            <Form onSubmit = {this.handleSubmit}>
            {/* <Form.Group controlId="formMessage"> */}
              <div style = {{display: 'flex', alignItems: 'flex-end', position: 'absolute', bottom: '2%', width: '100'}}>
                <Form.Control
                type="text"
                placeholder="Send a message..."
                onChange={this.handleReadText}
                ref = {this.myQuestion}
                style={{ fontSize: '24px', marginLeft: '-15%', marginRight: '0px', width: '1400px'}}
                />
                <Button type = "submit" style = {{marginLeft: '5px', height: '50px'}} disabled={this.state.algo == 0}> send </Button>
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