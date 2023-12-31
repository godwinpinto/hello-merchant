CREATE DATABASE `UPN_DB`;

USE DATABASE `UPN_DB`;

-- UPN_DB.UPN_SQUAREUP_MAPPING definition

CREATE TABLE `UPN_SQUAREUP_MAPPING` (
  `USM_ROW_ID` char(20) NOT NULL,
  `UUM_ROW_ID` char(20) NOT NULL,
  `URL_UUID` varchar(500) NOT NULL,
  `ACTIVE` char(1) NOT NULL DEFAULT 'Y',
  `CREATED_DT` timestamp NOT NULL,
  `CREATED_BY` varchar(50) NOT NULL,
  `UPDATED_DT` timestamp NOT NULL,
  `UPDATED_BY` varchar(100) NOT NULL,
  PRIMARY KEY (`USM_ROW_ID`)
);

-- UPN_DB.UPN_TRANSACTION_LOG definition

CREATE TABLE `UPN_TRANSACTION_LOG` (
  `UTL_ROW_ID` char(20) NOT NULL,
  `UUM_ROW_ID` char(20) NOT NULL,
  `LOG_DATA` varchar(8000) NOT NULL,
  PRIMARY KEY (`UTL_ROW_ID`)
);

-- UPN_DB.UPN_TRANSACTION_MASTER definition

CREATE TABLE `UPN_TRANSACTION_MASTER` (
  `UTM_ROW_ID` char(20) NOT NULL,
  `UUM_ROW_ID` char(20) NOT NULL,
  `AMOUNT` varchar(10) NOT NULL,
  `CURRENCY` varchar(5) NOT NULL,
  `CHANNEL` varchar(10) NOT NULL,
  `ACTIVE` char(1) NOT NULL DEFAULT 'Y',
  `CREATED_DT` timestamp NOT NULL,
  `TRANSACTION_ID` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`UTM_ROW_ID`) /*T![clustered_index] CLUSTERED */
);

-- UPN_DB.UPN_USER_MASTER definition

CREATE TABLE `UPN_USER_MASTER` (
  `UUM_ROW_ID` char(20) NOT NULL,
  `USER_ID` varchar(35) NOT NULL,
  `ROLE` char(1) NOT NULL,
  `ACTIVE` char(1) NOT NULL DEFAULT 'Y',
  `CREATED_DT` timestamp NOT NULL,
  `CREATED_BY` varchar(50) NOT NULL,
  `UPDATED_DT` timestamp NOT NULL,
  `UPDATED_BY` varchar(100) NOT NULL,
  PRIMARY KEY (`UUM_ROW_ID`) /*T![clustered_index] CLUSTERED */,
  KEY `UPN_USER_MASTER_USER_ID` (`USER_ID`)
);

-- UPN_DB.UPN_XRPL_MAPPING definition

CREATE TABLE `UPN_XRPL_MAPPING` (
  `UXM_ROW_ID` char(20) NOT NULL,
  `UUM_ROW_ID` char(20) NOT NULL,
  `XRPL_AC_NO` varchar(500) NOT NULL,
  `ACTIVE` char(1) NOT NULL DEFAULT 'Y',
  `CREATED_DT` timestamp NOT NULL,
  `CREATED_BY` varchar(50) NOT NULL,
  `UPDATED_DT` timestamp NOT NULL,
  `UPDATED_BY` varchar(100) NOT NULL,
  PRIMARY KEY (`UXM_ROW_ID`) /*T![clustered_index] CLUSTERED */
);

