import React from 'react';
import {CDBSidebar,
    CDBSidebarContent,
    CDBSidebarFooter,
    CDBSidebarHeader,
    CDBSidebarMenu,
    CDBSidebarMenuItem,} from 'cdbreact';
import { NavLink } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'mdb-react-ui-kit/dist/css/mdb.min.css';
import { Planet } from 'react-kawaii';
import styles from './SideBar.modules.css'
import RSwitch from "react-switch";
// import MainGPT from '../../pages/MainGPT/MainGPT';



class SideBar extends React.Component {  
    constructor(props){
      super(props);
      this.state = {
      isKMP : false,
      isBM : false
        };
    }

    setKMP = (e) =>{
      this.setState({
        isKMP : !this.state.isKMP
      })
      this.props.handleAlgoChange({target: {value: true}})
    }

    setBM = (e) =>{
      this.setState({
        isBM : !this.state.isBM
      })
      this.props.handleAlgoChange({target:{value : false}})
    }

    render() { 
        return(
                <div style={{ display: 'fixed', height: '100vh'}}>
                  <CDBSidebar textColor="#fff" backgroundColor="#B2A4FF">
                    <CDBSidebarHeader prefix={<i className="fa fa-bars fa-large"></i>}>
                      <a href="/" className="text-decoration-none" style={{ color: 'inherit' ,display: 'flex',alignItems: 'center', padding: '20px 0px',}}>
                      <Planet size={80} mood="blissful" color="#FDA7DC" />&nbsp;&nbsp;&nbsp; Chatty-Loli
                      </a>
                    </CDBSidebarHeader>
            
                    <CDBSidebarContent className="sidebar-content" style = {{display: 'flex', flexDirection: 'column', alignItems: 'center',marginTop: '-20%', fontSize: '24px'}}>
                      <CDBSidebarMenu>                       
                          <CDBSidebarMenuItem icon="history" >History</CDBSidebarMenuItem>
                      </CDBSidebarMenu>
                    </CDBSidebarContent>
            
                    <CDBSidebarFooter style={{ textAlign: 'center' , alignItems: 'center' }}>
                      <div
                        style={{
                          position : 'relative',
                          bottom : '20%',
                          display: 'inline-block',
                          padding: '40px 20px',
                        
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