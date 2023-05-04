import React from 'react';
import {CDBSidebar,
    CDBSidebarContent,
    CDBSidebarFooter,
    CDBSidebarHeader,
    CDBSidebarMenu,
    CDBSidebarMenuItem} from 'cdbreact';
import { NavLink } from 'react-router-dom';
import { Button } from "react-bootstrap";
import 'bootstrap/dist/css/bootstrap.min.css';
import 'mdb-react-ui-kit/dist/css/mdb.min.css';
import { Planet } from 'react-kawaii';
import styles from './SideBar.module.css'
import RSwitch from "react-switch";
import axios from "axios"
// import MainGPT from '../../pages/MainGPT/MainGPT';



class SideBar extends React.Component {  
    constructor(props){
      super(props);
      this.state = {
      isKMP : false,
      isBM : false,
      showHistory: true,
      history : [],
        };
    }

    setKMP = (e) =>{
      this.setState({
        isKMP : !this.state.isKMP
      })
      if(!this.state.isKMP)this.props.handleAlgoChange({target: {value: true}});
      else this.props.handleAlgoChange({target: {value: this.state.isKMP}});
    }

    setBM = (e) =>{
      this.setState({
        isBM : !this.state.isBM
      })
      if(!this.state.isBM)this.props.handleAlgoChange({target: {value: false}});
      else this.props.handleAlgoChange({target: {value: this.state.isKMP}});
    }
    componentDidMount(){
      axios.get("http://localhost:8000/api/session")
        .then(response =>{
          const temp = [];
          let n = response.data.length-1;
          let count = 0;
          while(n>=0 && count < 9){
              count++;
              console.log(temp[n])
              temp.push(response.data[n].Session)
              n--;
          }

          this.setState({history: temp})
        })
        .catch(error => {
          console.log(error)
        })
    }
    
    addHistory = async(e) =>{
      e.preventDefault()
      let temp = this.state.history
      let now = Date.now()
      console.log(now)
      const data = {
        Session : now
      }
      await axios.post("http://localhost:8000/api/session", data)
      
      if(temp.length < 10){
        temp.push(now)
      }
      else{
        
        let oldest = -1
        for(let i = 0; i < this.state.history.length; i++){
          if(i === 0){
            oldest = 0
          }else{
            if(this.state.history[i] < this.state.history[oldest]){
              oldest = i
            }
          }
          
        }
        temp[oldest] = now
        window.alert('Data History ' + this.state.history[oldest] + ' telah dioverwrite');
      }
      this.setState({history: temp})

      this.props.handleNewSession({target: {value: now}})
    }
    chooseHistory(historyCode){
      this.props.handleHistory({target: {value: historyCode}})
    }



    render_history() {
      let historyItems = [];
      for (let i = 0; i < this.state.history.length; i++) {
        let historyCode = "History " + ('0' + this.state.history[i]).slice(-2); // format nomor menjadi H01, H02, dst.

        historyItems.push(<li color='purple' id = {this.state.history[i]}  onClick={() => this.chooseHistory(this.state.history[i])}>{historyCode}</li>);
      }
      return (
        <div className={styles.history}>
          <ul>
            {historyItems}
          </ul>
        </div>
      );
    }
    
    render() { 
        //Todo: get history id from database:

        return(
                <div style={{ display: 'fixed', height: '100vh'}}>
                  <CDBSidebar textColor="#fff" backgroundColor="#B2A4FF">
                    <CDBSidebarHeader prefix={<i className="fa fa-bars fa-large"></i>}>
                      <a href="/" className="text-decoration-none" style={{ color: 'inherit' ,display: 'flex',alignItems: 'center', padding: '20px 0px',}} >
                      <Planet size={80} mood="blissful" color="#FDA7DC" />&nbsp;&nbsp;&nbsp; Chatty-Loli
                      </a>
                    </CDBSidebarHeader>
            
                    <CDBSidebarContent className="sidebar-content" style = {{display: 'flex', flexDirection: 'column', alignItems: 'center',marginTop: '-20%', fontSize: '24px'}}>
                      <CDBSidebarMenu>                       
                          <CDBSidebarMenuItem icon="history" >History</CDBSidebarMenuItem>
                      </CDBSidebarMenu>
                      <CDBSidebarMenu style ={{marginTop:'170px'}}> 
                         <CDBSidebarMenuItem >{this.render_history()}</CDBSidebarMenuItem>
                      </CDBSidebarMenu>
                      <Button key={Math.random()} type = "submit" style = {{height: '40px', width:'100px', marginTop: '200px'}} onClick = {this.addHistory} > +Add </Button>
                    </CDBSidebarContent>

                    
                    <CDBSidebarFooter style={{ textAlign: 'center' , alignItems: 'center' }}>
                      <div
                        style={{
                          position : 'relative',
                          bottom : '0%',
                          display: 'inline-block',
                          padding: '20px 20px',
                        
                        }}
                      >
                      <label >
                        <span><p>KMP</p></span>
                        <RSwitch onChange={this.setKMP} checked={this.state.isKMP} 
                                onColor = '#8d4ca7' checkedIcon={false}uncheckedIcon={false} disabled = {this.state.isBM}/>
                      </label>
                      <p></p>
                      <label>
                        <span><p>Boyer-Moore</p></span>
                        <RSwitch onChange={this.setBM} checked={this.state.isBM} 
                                onColor = '#8d4ca7' checkedIcon={false}uncheckedIcon={false} disabled = {this.state.isKMP}/>
                      </label>
                      </div>
                    </CDBSidebarFooter>
                  </CDBSidebar>
                  
                </div>
              );
     }
}
     
     
  export default SideBar;