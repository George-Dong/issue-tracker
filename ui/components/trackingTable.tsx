import * as React from "react";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import Paper from "@mui/material/Paper";
import { ClosedIssueInfo, CloserPRInfo } from "./types";

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString(undefined, {
    year: "numeric",
    month: "numeric",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  });
};

export default function BasicTable(props: { infos: ClosedIssueInfo[] }) {
  const infos = props.infos;
  return (
    <TableContainer component={Paper}>
      <Table
        stickyHeader
        sx={{ minWidth: 650, maxHeight: "max-content" }}
        size="small"
        aria-label="simple table"
      >
        <TableHead>
          <TableRow>
            <TableCell>Issue</TableCell>
            <TableCell align="right">Closed At</TableCell>
            <TableCell align="right">Severity</TableCell>
            <TableCell align="right">Affected Version</TableCell>
            <TableCell align="right">Closed By</TableCell>
            <TableCell align="right">Cherry Picked To</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {infos.map((row) => (
            <TableRow
              key={row.Number}
              sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
            >
              <TableCell component="th" scope="row">
                <a href={row.Url}>
                  #{row.Number} {row.Title}
                </a>
              </TableCell>
              <TableCell align="right">{formatDate(row.ClosedAt)}</TableCell>
              <TableCell align="right">{row.Severity}</TableCell>
              <TableCell align="right">
                {row.AffectedVersions !== null &&
                  row.AffectedVersions.map((v) => (
                    <>
                      <code>{v}</code>
                      <br />
                    </>
                  ))}
              </TableCell>
              <TableCell align="right">
                {row.ClosedByPR === null ? (
                  "manually closed"
                ) : (
                  <a href={row.ClosedByPR.Url}>#{row.ClosedByPR.Number}</a>
                )}
              </TableCell>
              <TableCell align="right">
                {row.CloserCherryPicked !== null &&
                  row.CloserCherryPicked.map((cp: CloserPRInfo) => {
                    return (
                      <div>
                        {cp.MergeTarget}
                        {""}
                        <a href={cp.Url} key={cp.Number}>
                          #{cp.Number}
                        </a>
                        {""}
                        {cp.State}
                        <br />
                      </div>
                    );
                  })}
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
