import React, { useState } from 'react';
import { IconContext } from 'react-icons';
import { IoIosCalendar } from 'react-icons/io';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Badge from 'react-bootstrap/Badge';
import Card from 'react-bootstrap/Card';
import dayjs from 'dayjs';
import useAsync from '../hooks/useAsync';
import API from '../service/APIService';

interface SubscriptionCardProps {
  userid: string;
}

const SubscriptionCard: React.FC<SubscriptionCardProps> = ({ userid }) => {
  const { data, loading, error } = useAsync(() => API.GetUserPlanByUserID(userid));

  if (loading || data === undefined) {
    return <div></div>;
  }

  const { userplan } = data;
  let status = userplan.active ? 'Active' : 'Disabled';
  return (
    <Card className="dash-card">
      <Card.Body>
        <Card.Title>
          Subscription
          {userplan.plan_type == 'FREE' ? (
            <Badge className="sub-status" variant="info">
              Trial
            </Badge>
          ) : (
            <Badge className="sub-status" variant={status == 'Active' ? 'success' : 'warning'}>
              {status}
            </Badge>
          )}
        </Card.Title>
        <hr></hr>
        <div className="sub-card">
          <Row>
            <Col>
              <label htmlFor="plan_name">Name</label>
              <div id="plan_name">{userplan.plan_name}</div>
            </Col>
          </Row>
          <Row>
            <Col>
              <label htmlFor="startdate">
                <IconContext.Provider value={{ color: 'black' }}>
                  <div className="sub-label">
                    <IoIosCalendar />
                    <div>Start</div>
                  </div>
                </IconContext.Provider>
              </label>
              <div id="startdate">
                {dayjs(userplan.start_date)
                  .format('DD-MM-YYYY')
                  .toString()}
              </div>
            </Col>
            <Col>
              <label htmlFor="enddate">
                <IconContext.Provider value={{ color: 'black' }}>
                  <div className="sub-label">
                    <IoIosCalendar />
                    <div>End</div>
                  </div>
                </IconContext.Provider>
              </label>
              <div id="enddate">
                {dayjs(userplan.expiry_date)
                  .format('DD-MM-YYYY')
                  .toString()}
              </div>
            </Col>
          </Row>
        </div>
      </Card.Body>
    </Card>
  );
};

export default SubscriptionCard;
