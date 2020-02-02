import React, { useState, useEffect } from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import Auth from '../service/Auth';
import Link from 'next/link';

const handleLogout = async () => {
  Auth.Logout();
};

const HeaderBrand = () => {
  return (
    <Navbar.Brand>
      <Link href="/">
        <a>
          <h2 className="branding">
            ÉireVPN
            <Badge className="badge-custom" variant="danger">
              Beta
            </Badge>
          </h2>
        </a>
      </Link>
      <style jsx>{`
        a {
          text-decoration: none;
        }
        a:hover {
          opacity: 0.6;
        }
      `}</style>
    </Navbar.Brand>
  );
};

interface NavLinkProps {
  href: string;
  text: string;
}

const NavLink: React.FC<NavLinkProps> = ({ href, text }) => {
  return (
    <Navbar.Text>
      <Link href={href}>
        <a>
          <div className="header-link">{text}</div>
        </a>
      </Link>
      <style jsx>{`
        a {
          text-decoration: none;
        }
        a:hover {
          opacity: 0.6;
        }
      `}</style>
    </Navbar.Text>
  );
};

const HeaderAdmin = () => {
  const [loginStatus, setLoginStatus] = useState('Login');
  useEffect(() => {
    setLoginStatus(Auth.IsLoggedIn() ? 'Logout' : 'Login');
  });

  return (
    <Navbar>
      <HeaderBrand />
      <Navbar.Toggle />
      <Navbar.Collapse className="justify-content-end">
        <Navbar.Text onClick={handleLogout}>
          <a href="#">{loginStatus}</a>
        </Navbar.Text>
      </Navbar.Collapse>
    </Navbar>
  );
};

const HeaderUser = () => {
  const [loginStatus, setLoginStatus] = useState('Login');
  useEffect(() => {
    setLoginStatus(Auth.IsLoggedIn() ? 'Logout' : 'Login');
  });

  return (
    <Navbar className="header" expand="lg">
      <HeaderBrand />
      <Navbar.Toggle className="navbar-toggler-custom" />
      <Navbar.Collapse className="justify-content-end">
        <NavLink href="/about" text="About" />
        <NavLink href="/products" text="Products" />
        <NavLink href="/login" text={loginStatus} />
        <Navbar.Text>
          <Button className="btn-landing sm">Get Started</Button>
        </Navbar.Text>
      </Navbar.Collapse>
    </Navbar>
  );
};

const HeaderLogin = () => {
  const [loginStatus, setLoginStatus] = useState('Login');
  useEffect(() => {
    setLoginStatus(Auth.IsLoggedIn() ? 'Logout' : 'Login');
  });

  return (
    <Navbar expand="lg">
      <HeaderBrand />
      <Navbar.Toggle />
      <Navbar.Collapse className="justify-content-end"></Navbar.Collapse>
    </Navbar>
  );
};

export { HeaderAdmin, HeaderUser, HeaderLogin };
