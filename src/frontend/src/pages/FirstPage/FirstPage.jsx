import React, {Component,Fragment} from 'react';
import styles from './FirstPage.module.css';
import 'bootstrap/dist/css/bootstrap.css';
import MyImage from '../../attributes/about_1.png'
import { Cat} from 'react-kawaii';



class FirstPage extends React.Component {  
    constructor(props){
      super(props);
      this.state = {
      };
      
    }
    

    render() { 
      return (
    <body>

      <div className={styles.head}>
        <div className={styles.title}>
          <span><Cat size={80} mood="lovestruck" color="#FDA7DC" />CHATTY-LOLI</span>
          <span >
              <a href="/about">about</a>
          </span>
        </div>
      </div>

     
      <div className={styles.intro}>
        <div className={styles.content}>
        
          <span>
            Answer your questions easily!
          </span>
          <p>
          Chatty-Loli is a web based chat ai that will help you answer many questions easily. 
          It’s also smart enough to do simple math and guess the day of a date.
          </p>
        </div>
        <img src={MyImage} alt="Intro" />
      </div>

      <div className={styles.button}>
        <a href="/MainGPT">
          <button type="button" class="btn btn-outline-primary btn-lg">Get Started!</button>
        </a>
      </div>
     

      
      
    </body>       
      );
    }
}
     
     
  export default FirstPage;