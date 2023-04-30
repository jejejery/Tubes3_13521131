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
        algo : true,
        answer: ""
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
        Input : this.state.text,
        Algorithm: this.state.algo,
      };
      await axios.post("http://localhost:8000/api/input", data)
        .then(response => {
          this.setState({answer : response.data.result})

          console.log(data.Input);
          const pattern1 = /tambahkan pertanyaan (.+) dengan jawaban (.+)/i;
          const pattern2 = /hapus pertanyaan (.+)/i;
          const matches1 = data.Input.match(pattern1);
          const matches2 = data.Input.match(pattern2);
          if(matches1){
            // todo: question is available
            const newQna = {
              Question: matches1[1],
              Answer : matches1[2],
            }
              axios.post("http://localhost:8000/api/qna", newQna);
          }
          else if(matches2){
            const deleteQna = {
              Question: matches2[1],
              Answer: "",
            }
            axios.delete("http://localhost:8000/api/qna", {data: deleteQna});
          }
   
          }
        )
        .catch(error =>{
          console.log(error);
        })
        this.setState({text: ""})

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
            value = {this.state.text}
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