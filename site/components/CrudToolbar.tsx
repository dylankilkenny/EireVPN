import React from 'react';
import ButtonMain from './ButtonMain';
import Button from 'react-bootstrap/Button';

interface CrudToolbarProps {
  HandleCreate: () => void;
}

const CrudToolbar: React.FC<CrudToolbarProps> = ({ HandleCreate }): JSX.Element => {
  const handleCreate = () => {
    HandleCreate();
  };
  return (
    <div className="crud-toolbar">
      <ButtonMain onClick={handleCreate} value="Create" />
    </div>
  );
};

export default CrudToolbar;
