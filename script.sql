USE [master]
GO
/****** Object:  Database [mainDb]    Script Date: 16.02.2022 11:26:04 ******/
CREATE DATABASE [mainDb]
 CONTAINMENT = NONE
 ON  PRIMARY 
( NAME = N'mainDb', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER\MSSQL\DATA\mainDb.mdf' , SIZE = 8192KB , MAXSIZE = UNLIMITED, FILEGROWTH = 65536KB )
 LOG ON 
( NAME = N'mainDb_log', FILENAME = N'C:\Program Files\Microsoft SQL Server\MSSQL15.MSSQLSERVER\MSSQL\DATA\mainDb_log.ldf' , SIZE = 8192KB , MAXSIZE = 2048GB , FILEGROWTH = 65536KB )
 WITH CATALOG_COLLATION = DATABASE_DEFAULT
GO
ALTER DATABASE [mainDb] SET COMPATIBILITY_LEVEL = 150
GO
IF (1 = FULLTEXTSERVICEPROPERTY('IsFullTextInstalled'))
begin
EXEC [mainDb].[dbo].[sp_fulltext_database] @action = 'enable'
end
GO
ALTER DATABASE [mainDb] SET ANSI_NULL_DEFAULT OFF 
GO
ALTER DATABASE [mainDb] SET ANSI_NULLS OFF 
GO
ALTER DATABASE [mainDb] SET ANSI_PADDING OFF 
GO
ALTER DATABASE [mainDb] SET ANSI_WARNINGS OFF 
GO
ALTER DATABASE [mainDb] SET ARITHABORT OFF 
GO
ALTER DATABASE [mainDb] SET AUTO_CLOSE OFF 
GO
ALTER DATABASE [mainDb] SET AUTO_SHRINK OFF 
GO
ALTER DATABASE [mainDb] SET AUTO_UPDATE_STATISTICS ON 
GO
ALTER DATABASE [mainDb] SET CURSOR_CLOSE_ON_COMMIT OFF 
GO
ALTER DATABASE [mainDb] SET CURSOR_DEFAULT  GLOBAL 
GO
ALTER DATABASE [mainDb] SET CONCAT_NULL_YIELDS_NULL OFF 
GO
ALTER DATABASE [mainDb] SET NUMERIC_ROUNDABORT OFF 
GO
ALTER DATABASE [mainDb] SET QUOTED_IDENTIFIER OFF 
GO
ALTER DATABASE [mainDb] SET RECURSIVE_TRIGGERS OFF 
GO
ALTER DATABASE [mainDb] SET  DISABLE_BROKER 
GO
ALTER DATABASE [mainDb] SET AUTO_UPDATE_STATISTICS_ASYNC OFF 
GO
ALTER DATABASE [mainDb] SET DATE_CORRELATION_OPTIMIZATION OFF 
GO
ALTER DATABASE [mainDb] SET TRUSTWORTHY OFF 
GO
ALTER DATABASE [mainDb] SET ALLOW_SNAPSHOT_ISOLATION OFF 
GO
ALTER DATABASE [mainDb] SET PARAMETERIZATION SIMPLE 
GO
ALTER DATABASE [mainDb] SET READ_COMMITTED_SNAPSHOT OFF 
GO
ALTER DATABASE [mainDb] SET HONOR_BROKER_PRIORITY OFF 
GO
ALTER DATABASE [mainDb] SET RECOVERY FULL 
GO
ALTER DATABASE [mainDb] SET  MULTI_USER 
GO
ALTER DATABASE [mainDb] SET PAGE_VERIFY CHECKSUM  
GO
ALTER DATABASE [mainDb] SET DB_CHAINING OFF 
GO
ALTER DATABASE [mainDb] SET FILESTREAM( NON_TRANSACTED_ACCESS = OFF ) 
GO
ALTER DATABASE [mainDb] SET TARGET_RECOVERY_TIME = 60 SECONDS 
GO
ALTER DATABASE [mainDb] SET DELAYED_DURABILITY = DISABLED 
GO
ALTER DATABASE [mainDb] SET ACCELERATED_DATABASE_RECOVERY = OFF  
GO
EXEC sys.sp_db_vardecimal_storage_format N'mainDb', N'ON'
GO
ALTER DATABASE [mainDb] SET QUERY_STORE = OFF
GO
USE [mainDb]
GO
/****** Object:  Table [dbo].[Project&Users]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Project&Users](
	[Id] [uniqueidentifier] NOT NULL,
	[UserId] [uniqueidentifier] NOT NULL,
	[ProjectId] [uniqueidentifier] NOT NULL,
 CONSTRAINT [PK_Project&Users] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Projects]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Projects](
	[Id] [uniqueidentifier] NOT NULL,
	[ParentId] [uniqueidentifier] NULL,
	[ProjectName] [varchar](50) NOT NULL,
	[ProjectDescription] [varchar](500) NULL,
	[CreationDate] [date] NOT NULL,
	[OwnerId] [uniqueidentifier] NOT NULL,
 CONSTRAINT [PK_Projects] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Roles]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Roles](
	[Id] [uniqueidentifier] NOT NULL,
	[RoleName] [varchar](50) NOT NULL,
 CONSTRAINT [PK_Roles] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Tags]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Tags](
	[Id] [uniqueidentifier] NOT NULL,
	[TaskName] [varchar](50) NOT NULL,
	[TaskColor] [varchar](6) NOT NULL,
 CONSTRAINT [PK_Tags] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Task&Tags]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Task&Tags](
	[Id] [uniqueidentifier] NOT NULL,
	[TaskId] [uniqueidentifier] NOT NULL,
	[TagId] [uniqueidentifier] NOT NULL,
 CONSTRAINT [PK_Task&Tags] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Tasks]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Tasks](
	[Id] [uniqueidentifier] NOT NULL,
	[ProjectId] [uniqueidentifier] NOT NULL,
	[TaskTitle] [varchar](50) NOT NULL,
	[TaskDescription] [varchar](250) NULL,
 CONSTRAINT [PK_Tasks] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
/****** Object:  Table [dbo].[Users]    Script Date: 16.02.2022 11:26:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[Users](
	[Id] [uniqueidentifier] NOT NULL,
	[Name] [varchar](50) NOT NULL,
	[Surname] [varchar](50) NOT NULL,
	[Nickname] [varchar](50) NOT NULL,
	[RoleId] [uniqueidentifier] NOT NULL,
	[Password] [varchar](50) NOT NULL,
 CONSTRAINT [PK_Users] PRIMARY KEY CLUSTERED 
(
	[Id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON, OPTIMIZE_FOR_SEQUENTIAL_KEY = OFF) ON [PRIMARY]
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[Project&Users]  WITH CHECK ADD  CONSTRAINT [ProjectId_to_Project&UsersProjectId] FOREIGN KEY([ProjectId])
REFERENCES [dbo].[Projects] ([Id])
GO
ALTER TABLE [dbo].[Project&Users] CHECK CONSTRAINT [ProjectId_to_Project&UsersProjectId]
GO
ALTER TABLE [dbo].[Project&Users]  WITH CHECK ADD  CONSTRAINT [UserId_to_Project&UsersUserId] FOREIGN KEY([UserId])
REFERENCES [dbo].[Users] ([Id])
GO
ALTER TABLE [dbo].[Project&Users] CHECK CONSTRAINT [UserId_to_Project&UsersUserId]
GO
ALTER TABLE [dbo].[Projects]  WITH CHECK ADD  CONSTRAINT [ProjectId_to_ParentId] FOREIGN KEY([ParentId])
REFERENCES [dbo].[Projects] ([Id])
GO
ALTER TABLE [dbo].[Projects] CHECK CONSTRAINT [ProjectId_to_ParentId]
GO
ALTER TABLE [dbo].[Projects]  WITH CHECK ADD  CONSTRAINT [UserId_to_ProjectOwnerId] FOREIGN KEY([OwnerId])
REFERENCES [dbo].[Users] ([Id])
GO
ALTER TABLE [dbo].[Projects] CHECK CONSTRAINT [UserId_to_ProjectOwnerId]
GO
ALTER TABLE [dbo].[Task&Tags]  WITH CHECK ADD  CONSTRAINT [TagId_to_Task&TagsTagId] FOREIGN KEY([TagId])
REFERENCES [dbo].[Tags] ([Id])
GO
ALTER TABLE [dbo].[Task&Tags] CHECK CONSTRAINT [TagId_to_Task&TagsTagId]
GO
ALTER TABLE [dbo].[Task&Tags]  WITH CHECK ADD  CONSTRAINT [TaskId_to_Task&TagsTaskId] FOREIGN KEY([TaskId])
REFERENCES [dbo].[Tasks] ([Id])
GO
ALTER TABLE [dbo].[Task&Tags] CHECK CONSTRAINT [TaskId_to_Task&TagsTaskId]
GO
ALTER TABLE [dbo].[Tasks]  WITH CHECK ADD  CONSTRAINT [ProjectId_to_TaskProjectId] FOREIGN KEY([ProjectId])
REFERENCES [dbo].[Projects] ([Id])
GO
ALTER TABLE [dbo].[Tasks] CHECK CONSTRAINT [ProjectId_to_TaskProjectId]
GO
ALTER TABLE [dbo].[Users]  WITH CHECK ADD  CONSTRAINT [RoleId_to_UsersRoleId] FOREIGN KEY([RoleId])
REFERENCES [dbo].[Roles] ([Id])
GO
ALTER TABLE [dbo].[Users] CHECK CONSTRAINT [RoleId_to_UsersRoleId]
GO
USE [master]
GO
ALTER DATABASE [mainDb] SET  READ_WRITE 
GO
