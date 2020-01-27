import React from 'react';
import Link from 'next/link';
import ButtonToolbar from 'react-bootstrap/ButtonToolbar';
import Button from 'react-bootstrap/Button';

interface CrudToolbarProps {
  HandleCreate: () => void;
}

const CrudToolbar: React.FC<CrudToolbarProps> = ({
  HandleCreate
}): JSX.Element => {
  const handleCreate = () => {
    HandleCreate();
  };
  return (
    <div className="button-toolbar">
      <Button onClick={handleCreate} variant="outline-secondary">
        Create
      </Button>
    </div>
  );
};

export default CrudToolbar;
