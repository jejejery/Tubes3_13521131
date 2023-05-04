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
        algo : null,//true : kmp, false: bm
        qaBlocks : [],
        formRows: 1,
        answer: "",
        startSession: Date.now()
      };
      this.handleKeyPress = this.handleKeyPress.bind(this)
      this.setState = this.setState.bind(this)

    }
    handleAlgoChange = (event) =>{
      this.setState({algo: event.target.value})
    }
    handleReadText = (event) =>{
      this.myQuestion.current.value = event.target.value;
    }
    handleNewSession = (event) => {
      this.setState({startSession: event.target.value})
      this.setState({qaBlocks : []})
  
    }
    handleHistory = async(event) => {
      await axios.get(`http://localhost:8000/api/input/toShow?Session=${event.target.value}`)
      .then(response => {
        console.log(response.data)
        let temp = []
        for(let i = 0; i < response.data.length; i++){
          temp.push(this.render_q_block(response.data[i].InputText,response.data[i].Answer))
        }
        this.setState({qaBlocks: temp})
        }
      )
      .catch(error =>{
        console.log(error);
      })
      
    }
    handleSubmit = async (event) => {
      event.preventDefault()
      console.log(this.myQuestion.current.value)
      console.log(this.state.algo)
      console.log(this.state.startSession)

      const data = {
        Session: this.state.startSession,
        Input :this.myQuestion.current.value,
        Algorithm: this.state.algo
        
      };
      await axios.post("http://localhost:8000/api/input", data)
      .then(response => {
        let temp = this.state.qaBlocks
        temp.push(this.render_q_block(this.myQuestion.current.value,response.data.Answer))
        this.setState({qaBlocks: temp})
        this.setState({answer : response.data.Answer})
 
        }
      )
      .catch(error =>{
        console.log(error);
      })

      //insert into array

      
      
      //try-catch
      
      

        
      //reset the question
      this.myQuestion.current.value = ""
    }
    handleKeyPress(event) {
      if (event.key == 'Enter') {
        if(this.state.formRows < 3) this.setState({ formRows: this.state.formRows+1 });
      }
    }
    
    render_q_block(question, answer){
      const questionLineCount = (question.match(/\n/g) || []).length + 1;
      const answerLineCount = (answer.match(/\n/g) || []).length + 1;
    
      const qblockHeight = 150 + (questionLineCount - 1) * 20; // 20 is the line height
      const ablockHeight = 100 + (answerLineCount - 1) * 20;
    
      return (
        <div>
          <div class={styles.qblock} style={{height: qblockHeight + 'px'}}>
            <div>
              <SpeechBubble size={80} mood="excited" color="rgb(90, 88, 177)" /> 
              <span dangerouslySetInnerHTML={{__html: question.replace(/\n/g, '<br>') }} />
            </div>
          </div>
    
          <div class={styles.ablock} style={{height: ablockHeight + 'px'}}>
            <div>
              <SpeechBubble size={80} mood="happy" color="#7e4474" /> 
              <span dangerouslySetInnerHTML={{__html: answer.replace(/\n/g, '<br>') }} />
            </div>          
          </div>
        </div>
      );
    }

    render() { 
      //To Do in here: read recent qna from database to show em up in window


      return (
        <div className={styles.main}>
        <div style={styles.first}>
          <SideBar style={{ margin: '0 10px' }} handleAlgoChange= {this.handleAlgoChange} handleNewSession= {this.handleNewSession} handleHistory = {this.handleHistory}/>
          <div className={styles.chatbox}>
            {this.state.qaBlocks}
          </div>
          <div style={{  left: '30%', bottom: '3%', flex: 5, position: 'absolute', borderColor:'black'}}>
            <Form onSubmit = {this.handleSubmit}>
            <Form.Group className="mb-3" controlId="formMessage">
              <div style = {{display: 'flex', alignItems: 'flex-end', position: 'absolute', bottom: '2%', width: '100'}}>
                <Form.Control
                type="text"
                as="textarea"
                placeholder="Send a message..."
                onChange={this.handleReadText}
                ref = {this.myQuestion}
                onKeyPress={this.handleKeyPress}
                rows = {this.state.formRows}
                style={{ fontSize: '24px', marginLeft: '-15%', marginRight: '0px', width: '1400px',}}
                />
                <Button type = "submit" style = {{marginLeft: '5px', height: '50px'}} disabled={this.state.algo == null}> send </Button>
              </div>
            </Form.Group>
            </Form>
         </div>
       </div>
      </div>
      
            );
     }
}
     
     
  export default MainGPT;