import Pagination from 'react-bootstrap/Pagination';
import { useState } from 'react';

interface PaginationProps {
  count: number;
  pageLimit: number;
  handlePagination: (page_number: number) => void;
}

const Pages: React.FC<PaginationProps> = ({ count, handlePagination, pageLimit }): JSX.Element => {
  const [activePage, setActivePage] = useState(1);
  const handleClick = (page: number) => {
    setActivePage(page);
    handlePagination(page);
  };
  const num_pages = Math.ceil(count / pageLimit);
  let items = [];
  for (let i = 1; i <= num_pages; i++) {
    items.push(
      <Pagination.Item onClick={() => handleClick(i)} key={i} active={i === activePage}>
        {i}
      </Pagination.Item>
    );
  }
  return <Pagination>{items}</Pagination>;
};

export default Pages;
