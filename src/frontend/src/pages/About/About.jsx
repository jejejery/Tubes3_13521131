import React, {Component,Fragment} from 'react';
import styles from './About.module.css';
import 'bootstrap/dist/css/bootstrap.css';
import { Cat} from 'react-kawaii';



class About extends React.Component {  
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
          <span><Cat size={80} mood="lovestruck" color="#FDA7DC" />About</span>
        </div>
      </div>

     
      <div className={styles.intro}>
        <div className={styles.content}>
    
          <p>
          Website ini dibuat oleh Dida, Gita, dan Jerry hehe :)
          </p>
        </div>
      </div>

     

      
      
    </body>       
      );
    }
}
     
     
  export default About;