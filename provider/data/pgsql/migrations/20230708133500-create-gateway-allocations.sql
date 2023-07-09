CREATE TYPE gatewayGroupRunningStatus AS ENUM (
    'RUNNING', -- means all desired containers are running
    'PENDING', -- means some of the containers are still not running and pending runtime allocation
    'ERROR', -- means no containers are running and there is some issue in local
    'INITIATED' -- means the container group has just started and initiation is in progress
);

CREATE TABLE "gatewayAllocations" (
      id SERIAL PRIMARY KEY NOT NULL,
      "projectID" INTEGER NOT NULL,
      "containerGroupID" INTEGER NOT NULL,
      "runningStatus" gatewayGroupRunningStatus,
      "statusDescription" VARCHAR(255),
      "statusAdditionalDetailsJSON" json,
      "gatewayHTTPPortConfigs" json NOT NULL, -- json for maintaining port configurations { "ports": [{ "containerPort": 2425 }, { "exposedPathName": "/anyPath" }] }
      "gatewayURL" VARCHAR(255) NOT NULL,
      "createdAt" timestamp with time zone NOT NULL,
      "updatedAt" timestamp with time zone NOT NULL,
      CONSTRAINT fk_projects FOREIGN KEY("projectID") REFERENCES projects(id),
      CONSTRAINT fk_container_groups FOREIGN KEY("containerGroupID")
          REFERENCES "containerAllocations"(id),
      CONSTRAINT unique_project_id UNIQUE("projectID")
);

