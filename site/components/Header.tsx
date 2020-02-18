import React, { useState, useEffect } from 'react';
import Navbar from 'react-bootstrap/Navbar';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import Image from 'react-bootstrap/Image';
import Auth from '../service/Auth';
import Link from 'next/link';

const handleLogout = async () => {
  Auth.Logout();
};

const HeaderBase: React.FC<React.ReactNode> = ({ children }) => {
  return (
    <Navbar className="navbar-cust header" expand="lg">
      <HeaderBrand />
      <Navbar.Toggle className="navbar-toggler-custom" />
      <Navbar.Collapse className="justify-content-end">{children}</Navbar.Collapse>
    </Navbar>
  );
};

const HeaderDashboard = () => {
  const [loginStatus, setLoginStatus] = useState('Login');
  useEffect(() => {
    setLoginStatus(Auth.IsLoggedIn() ? 'Logout' : 'Login');
  });

  return (
    <HeaderBase>
      <Navbar.Text onClick={handleLogout}>
        <a href="#">
          <div className="header-link">{loginStatus}</div>
        </a>
      </Navbar.Text>
      <style jsx>{`
        a {
          text-decoration: none;
        }
        a:hover {
          opacity: 0.6;
        }
      `}</style>
    </HeaderBase>
  );
};

const HeaderUser = () => {
  const [loginStatus, setLoginStatus] = useState('Login');
  const [loginLink, setLoginLink] = useState('Login');
  useEffect(() => {
    setLoginStatus(Auth.IsLoggedIn() ? 'My Account' : 'Login');
    setLoginLink(Auth.IsLoggedIn() ? '/account' : '/login');
  });

  return (
    <HeaderBase>
      <NavLink href="/products" text="Products" />
      <NavLink href="/contact" text="Contact Support" />
      <NavLink href={loginLink} text={loginStatus} />
      <Navbar.Text>
        <Link href="/signup">
          <Button className="btn-landing sm">Try Free</Button>
        </Link>
      </Navbar.Text>
    </HeaderBase>
  );
};

const HeaderLogin = () => {
  return <HeaderBase />;
};

const HeaderBrand = () => {
  return (
    <Navbar.Brand>
      <Link href="/">
        <a>
          <Image className="branding-shield" fluid src="../static/images/shield.png" />
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

export { HeaderDashboard, HeaderUser, HeaderLogin };
