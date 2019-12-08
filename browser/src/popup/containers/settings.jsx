import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Button from 'react-bootstrap/Button';
import PropTypes from 'prop-types';

class Settings extends React.Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  componentWillMount() {}

  componentDidMount() {}

  render() {
    return (
      <Container>
        <Row>
          <Col />
          <Col>
            <Button onClick={this.props.logout} variant="secondary">
              logout
            </Button>
          </Col>
          <Col />
        </Row>
      </Container>
    );
  }
}

Settings.propTypes = {
  logout: PropTypes.func.isRequired,
};
export default Settings;
