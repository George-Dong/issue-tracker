
export interface CloserPRInfo {
    Repository: string;
    Number: number;
    Title: string;
    Url: string;
    State: string;
    MergeTarget: string;
    MergedAt: string;
    Commit: string;
    Refs: string[];
}

export interface ClosedIssueInfo {
    Number: number;
    Title: string;
    Url: string;
    ClosedAt: string;
    Severity: string;
    AffectedVersions: string[];
    ClosedByPR: CloserPRInfo;
    CloserCherryPicked: CloserPRInfo[];
};