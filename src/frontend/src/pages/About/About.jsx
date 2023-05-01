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

    superblock(nim,nama){
      return (
        <div className={styles.superblock}>
          
          <div className={styles.photo}></div>
          <div className={styles.name}>
            <div>
            <span>{nim}</span>
            <span>{nama}</span>
            
            </div>
          </div>
          
        </div>
      );
    };

    render() { 
      return (
    <body>

      <div className={styles.head}>
        <div className={styles.title}>
          <span><Cat size={80} mood="happy" color="#FDA7DC" /><a href="/" >Home</a></span>
        </div>
      </div>

     
      <div className={styles.intro}>
        <div className={styles.content}>  
          This website developed by: 
        </div>
        <div style={{display: "flex", marginLeft: "300px", marginTop: "70px"}}>
          <div style={{display: "inline-block"}}>{this.superblock("13521131")}</div><span style={{marginRight: "10px"}}></span>
          <div style={{display: "inline-block"}}>{this.superblock("13521156")}</div><span style={{marginRight: "10px"}}></span>
          <div style={{display: "inline-block"}}>{this.superblock("13521157")}</div>
        </div>

        
      </div>

      <div className={styles.github}>  
        Want to know our repository?<br/>
        Check Our Github Repository Below!<br/>
        <a href="https://github.com/jejejery/Tubes3_IF2211" target="_blank" rel="noopener noreferrer">
         <button className={styles.button}>Github Repository</button>
        </a>
        
      </div>

     

      
      
    </body>       
      );
    }
}
     
     
  export default About;