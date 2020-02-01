/* global document */
import React from 'react';
import ReactDOM from 'react-dom';
import Popup from './containers/popup';
import '../assets/css/index.css';

const Index = () => <Popup />;

ReactDOM.render(<Index />, document.getElementById('display-container'));
