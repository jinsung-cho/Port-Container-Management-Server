import PropTypes from 'prop-types';
import {
  Card,
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableContainer
} from '@mui/material';
import { Scrollbar } from 'src/components/scrollbar';

export const PreInformation = (props) => {
  const {
    items = [],
  } = props;


  return (
    <Card>
      <Scrollbar>
        <Box sx={{ minWidth: 800 }}>
          <TableContainer sx={{ maxHeight: 420 }}>
            <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>
                    Equipment No
                  </TableCell>
                  <TableCell>
                    Container No
                  </TableCell>
                  <TableCell>
                    Truck No
                  </TableCell>
                  <TableCell>
                    Type No
                  </TableCell>
                  <TableCell>Date</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {items.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell>{item.inspeqno}</TableCell>
                    <TableCell>{item.cntrno}</TableCell>
                    <TableCell>{item.truckno}</TableCell>
                    <TableCell>{item.typeNo ? item.typeNo : "local DB"}</TableCell>
                    <TableCell>{item.qdate}</TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Scrollbar>
    </Card>
  );
};

PreInformation.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number,
    inspEqNo: PropTypes.string,
    cntrNo: PropTypes.string,
    truckNo: PropTypes.string,
  })),
};
