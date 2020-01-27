import React, { useState, useEffect } from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Auth from '../service/Auth';

export default function Header() {
  const [loginStatus, setLoginStatus] = useState('Login');
  useEffect(() => {
    setLoginStatus(Auth.IsLoggedIn() ? 'Logout' : 'Login');
  });

  const handleLogout = async () => {
    Auth.Logout();
  };

  return (
    <Navbar>
      <Navbar.Brand>
        <h3>EireVPN</h3>
      </Navbar.Brand>
      <Navbar.Toggle />
      <Navbar.Collapse className="justify-content-end">
        <Navbar.Text onClick={handleLogout}>
          <a href="#">{loginStatus}</a>
        </Navbar.Text>
      </Navbar.Collapse>
    </Navbar>
  );
}
