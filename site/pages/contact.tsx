import React, { useState } from 'react';
import { LayoutContact } from '../components/Layout';
import Image from 'react-bootstrap/Image';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import ContactForm from '../components/user/forms/ContactForm';
import API from '../service/APIService';
import ErrorMessage from '../components/ErrorMessage';
import SuccessMessage from '../components/SuccessMessage';

export default function ContactPage(): JSX.Element {
  const [success, setSuccess] = useState();
  const [error, setError] = useState();
  async function HandleSend(body: string) {
    const res = await API.ContactSupport(body);
    if (res.status == 200) {
      setSuccess(true);
      setError(false);
    } else {
      setError(res);
      setSuccess(false);
    }
  }
  return (
    <LayoutContact>
      <div className="contact-cont">
        <ErrorMessage show={!!error} error={error} />
        <SuccessMessage
          show={!!success}
          message="Message sent successfully. If you have not recieved an email confirming your message has been recieved, please check your spam folder."
        />
        <Container>
          <Row>
            <Col xs={12} sm={12} md={6}>
              <div className="contact-img-cont">
                <h1 className="center">Having an issue? Let us know.</h1>
                <Image className="contact-image" fluid src="../static/images/mail.png" />
              </div>
            </Col>
            <Col xs={12} sm={12} md={6}>
              <ContactForm success={success} HandleSend={HandleSend} />
            </Col>
          </Row>
        </Container>
      </div>
    </LayoutContact>
  );
}
