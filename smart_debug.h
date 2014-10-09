/*
 * =====================================================================================
 *
 *       Filename:  smart_debug.h
 *
 *    Description:  Describes what prints to exist and what not after preprocessing
 *                  Allows control over code footprint and printing messages based on level of criticality
 *
 *        Version:  1.0
 *        Created:  10/10/2014 00:01:56
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Ravi Teja. K (rtr), k.teza1@gmail.com
 *   Organization:  
 *
 * =====================================================================================
 */
#include <syslog.h>

#define PRINT(LEVEL, args...) SYSLOG_##LEVEL(##args)

#ifdef SYS_EMERG
#define SYSLOG_EMERG(args...) syslog(LOG_EMERG, ##args)
#else
#define SYSLOG_EMERG(args...)



#ifdef SYS_ALERT
#define SYSLOG_ALERT(args...) syslog(LOG_ALERT, ##args)
#else
#define SYSLOG_ALERT(args...)
#endif


#ifdef SYS_CRIT
#define SYSLOG_CRIT(args...) syslog(LOG_CRIT, ##args)
#else
#define SYSLOG_CRIT(args...)
#endif

#ifdef SYS_ERR
#define SYSLOG_ERR(args...) syslog(LOG_ERR, ##args)
#else
#define SYSLOG_ERR(args...)
#endif

#ifdef SYS_WARNING
#define SYSLOG_WARNING(args...) syslog(LOG_WARNING, ##args)
#else
#define SYSLOG_WARNING(args...)
#endif

#ifdef SYS_NOTICE
#define SYSLOG_NOTICE(args...) syslog(LOG_NOTICE, ##args)
#else
#define SYSLOG_NOTICE(args...)
#endif

#ifdef SYS_INFO
#define SYSLOG_INFO(args...) syslog(LOG_INFO, ##args)
#else
#define SYSLOG_INFO(args...)
#endif


#ifdef SYS_DEBUG
#define SYSLOG_DEBUG(args...) syslog(LOG_DEBUG, ##args)
#else
#define SYSLOG_INFO(args...)
#endif
