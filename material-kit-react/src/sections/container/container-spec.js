import PropTypes from 'prop-types';
import { useState } from 'react';
import {
  Card,
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  TableContainer, Button, Dialog, DialogTitle, DialogContent, DialogActions
} from '@mui/material';
import { Scrollbar } from 'src/components/scrollbar';
import axiosInstance from "src/axiosInstance";


export const ContainerSpec = (props) => {
  const {
    items = [],
  } = props;

  const [selectedRow, setSelectedRow] = useState(null);
  const [remarkData, setRemarkData] = useState([]);

  const handleRowClick = (itemId) => {
    if (selectedRow === itemId) {
      setSelectedRow(null);
    } else {
      setSelectedRow(itemId);
      axiosInstance.get("/remarks?informationid=eq." + itemId.toString())
        .then((response) => setRemarkData(response.data))
        .catch((error) => console.error('Error fetching data: ', error));
    }
  }

  const [open, setOpen] = useState(false);
  const handleOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

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
                    Inspection No
                  </TableCell>
                  <TableCell>
                    Start Time
                  </TableCell>
                  <TableCell>
                    End Time
                  </TableCell>
                  <TableCell>
                    Package Match
                  </TableCell>
                  <TableCell>
                    Inspection Result CD
                  </TableCell>
                  <TableCell>
                    Detection Count
                  </TableCell>
                  <TableCell>
                    Fault CD
                  </TableCell>
                  <TableCell>
                    Result Image Dir
                  </TableCell>
                  <TableCell>Date</TableCell>

                </TableRow>
              </TableHead>
              <TableBody>
                {items.map((item) => (
                  <>
                    <TableRow key={item.id}
                      onClick={() => handleRowClick(item.id)}
                      sx={{
                        backgroundColor: selectedRow === item.id ? 'rgba(0, 0, 255, 0.1)' : 'inherit',
                        fontWeight: selectedRow === item.id ? 'bold' : 'normal'
                      }}>
                      <TableCell>{item.inspeqno}</TableCell>
                      <TableCell>{item.inspno}</TableCell>
                      <TableCell>{item.inspstarttime}</TableCell>
                      <TableCell>{item.inspendtime}</TableCell>
                      <TableCell>{item.pckmatch}</TableCell>
                      <TableCell>{item.insprsltcd}</TableCell>
                      <TableCell>{item.detectioncnt}</TableCell>
                      <TableCell>{item.faultcd}</TableCell>
                      <TableCell>{item.insprsltimgdir}</TableCell>
                      <TableCell>{item.qdate}</TableCell>
                    </TableRow >
                    {selectedRow === item.id && (
                      <TableRow>
                        <TableCell colSpan={9}>
                          <Table>
                            <TableHead>
                              <TableRow>
                                <TableCell>Insp Remark</TableCell>
                                <TableCell align="right">
                                  <Button variant="contained" color="primary" onClick={handleOpen}>
                                    +Add remark
                                  </Button>
                                  <Dialog open={open} onClose={handleClose}>
                                    <DialogTitle>Add Remark</DialogTitle>
                                    <DialogContent>
                                    </DialogContent>
                                    <DialogActions>
                                      <Button onClick={handleClose} color="primary">
                                        Cancel
                                      </Button>
                                      <Button onClick={handleClose} color="primary">
                                        Add
                                      </Button>
                                    </DialogActions>
                                  </Dialog>
                                </TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>
                              {remarkData.map((remark) => (
                                <TableRow key={remark.id}>
                                  <TableCell>{remark.inspremark}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableCell>
                      </TableRow>
                    )}
                  </>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Box>
      </Scrollbar>
    </Card >
  );
};

ContainerSpec.propTypes = {
  items: PropTypes.arrayOf(PropTypes.shape({
    id: PropTypes.number,
    inspEqNo: PropTypes.string,
    inspNo: PropTypes.string,
    inspStartTime: PropTypes.string,
    inspEndTime: PropTypes.string,
    pckMatch: PropTypes.string,
    inspRsltCD: PropTypes.string,
    detectionCnt: PropTypes.string,
    faultCD: PropTypes.string,
    inspRsltImgDir: PropTypes.string,
  })),
};