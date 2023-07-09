CREATE TYPE containerGroupRunningStatus AS ENUM (
    'RUNNING', -- means all desired containers are running
    'PENDING', -- means some of the containers are still not running and pending runtime allocation
    'ERROR', -- means no containers are running and there is some issue in local
    'INITIATED' -- means the container group has just started and initiation is in progress
);

CREATE TABLE "containerAllocations" (
      id SERIAL PRIMARY KEY NOT NULL,
      "groupName" VARCHAR(255) NOT NULL,
      "projectID" integer NOT NULL,
      "containerCount" integer NOT NULL,
      "exposedTrafficPortsConfig" json NOT NULL,    -- where traffic will flow
      "containerRegistryURL" VARCHAR(255) NOT NULL,
      "runningContainerCount" int NOT NULL DEFAULT 0,
      "runningStatus" containerGroupRunningStatus,
      "statusDescription" VARCHAR(255),
      "statusAdditionalDetailsJSON" json,
      "createdAt" timestamp with time zone NOT NULL,
      "updatedAt" timestamp with time zone NOT NULL,
      CONSTRAINT fk_projects FOREIGN KEY("projectID") REFERENCES projects(id)
);

