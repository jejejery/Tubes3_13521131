import React, { Component } from 'react';
import styles from './About.module.css';
import 'bootstrap/dist/css/bootstrap.css';
import { Cat } from 'react-kawaii';

import jerry from '../../attributes/jerry.jpg';
import dida from '../../attributes/dida.jpg';
import gita from '../../attributes/gita.jpg';

class About extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  superblock(nim, nama, photoUrl) {
    return (
      <div className={styles.superblock}>
        <div className={styles.photo}>
          <img src={photoUrl} alt={`${nama} profile picture`} />
        </div>
        <div className={styles.name}>
          <div>
            <span>{nim}</span>
            <span>{nama}</span>
          </div>
        </div>
      </div>
    );
  }

  render() {
    return (
      <body>
        <div className={styles.head}>
          <div className={styles.title}>
            <span>
              <Cat size={80} mood="happy" color="#FDA7DC" />
              <a href="/">Home</a>
            </span>
          </div>
        </div>

        <div className={styles.intro}>
          <div className={styles.content}>This website developed by:</div>
          <div
            style={{ display: 'flex', marginLeft: '300px', marginTop: '70px' }}
          >
            <div style={{ display: 'inline-block' }}>
              {this.superblock(
                '13521131',
                'Jeremya Dharmawan R',
                jerry
              )}
            </div>
            <span style={{ marginRight: '10px' }}></span>
            <div style={{ display: 'inline-block' }}>
              {this.superblock(
                '13521157',
                'Brigita Tri Carolina',
                gita
              )}
            </div>
            <span style={{ marginRight: '10px' }}></span>
            <div style={{ display: 'inline-block' }}>
              {this.superblock(
                '13521156',
                'Kandida Edgina G',
                dida
              )}
            </div>
          </div>
        </div>

        <div className={styles.github}>
          Want to know our repository?
          <br />
          Check Our Github Repository Below!
          <br />
          <a
            href="https://github.com/jejejery/Tubes3_IF2211"
            target="_blank"
            rel="noopener noreferrer"
          >
            <button className={styles.button}>Github Repository</button>
          </a>
        </div>
      </body>
    );
  }
}

export default About;