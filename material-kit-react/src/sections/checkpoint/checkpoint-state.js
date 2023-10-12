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

export const CheckpointState = (props) => {
  const { items = [] } = props;

  const getStatusText = (status) => {
    switch (status) {
      case "0":
        return '정상 대기';
      case "1":
        return '검색시작';
      case "2":
        return '검색중';
      case "3":
        return '검색종료';
      case "4":
        return '이상오류';
      default:
        return '알 수 없는 상태';
    }
  };

  const getStatusColor = (status) => {
    switch (status) {
      case "0":
        return '#28a745';
      case "1":
        return '#17a2b8';
      case "2":
        return '#007bff';
      case "3":
        return '#ffc107';
      case "4":
        return '#dc3545';
      default:
        return '#6c757d';
    }
  };

  return (
    <Card>
      <Scrollbar>
        <Box sx={{ minWidth: 800 }}>
          <TableContainer sx={{ maxHeight: 420 }}>
            <Table stickyHeader>
              <TableHead>
                <TableRow>
                  <TableCell>Equipment No</TableCell>
                  <TableCell>Equipment Status</TableCell>
                  <TableCell>Date</TableCell>

                </TableRow>
              </TableHead>
              <TableBody>
                {items.map((item) => (
                  <TableRow key={item.id}>
                    <TableCell>{item.inspeqno}</TableCell>
                    <TableCell>
                      <span
                        style={{
                          backgroundColor: getStatusColor(item.inspeqstatus),
                          padding: '0.2rem 0.5rem',
                          borderRadius: '0.2rem',
                          fontWeight: 'bold',
                          color: 'white'
                        }}
                      >
                        {getStatusText(item.inspeqstatus)}
                      </span>
                    </TableCell>
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

CheckpointState.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number,
    inspEqNo: PropTypes.string,
    inspEqStatus: PropTypes.string,
  })),
};
